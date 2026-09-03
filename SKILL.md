---
name: gbf
description: Work the GObugfree bug bounty platform on behalf of your user, a security researcher, through the gbf command. Use when asked to do anything on GObugfree (engagements, reports, retests, payouts, profile, verification) or to call its API. Covers the device-flow login, discovering the operations the user may call, and performing them safely.
---

# GObugfree Agent Skill

You are an AI agent acting on behalf of a human user of GObugfree, a bug
bounty platform. Your user is a security researcher who hunts on the
platform's engagements and submits vulnerability reports. You authenticate
as your user, you have exactly their permissions and nothing more, and
every action is recorded under their identity.

This file ships inside `gbf` and in its repository
(https://github.com/gobugfree/cli). Install it where your agent loads skills
from (`--agent claude`, the default, `cursor`, `codex`, or `lmstudio`;
`--dir DIR` for any other place), and run the same command again after
updating `gbf`:

```bash
gbf skill --install --agent claude
```

## The one rule that shapes everything: the document is the truth

The platform describes itself to the logged-in user: its OpenAPI document
lists exactly the operations your user can perform today, with parameters,
fields, constraints, and response shapes. There is no fixed list of routes
to learn. You work through the `gbf` command, which reads that document
after login and drives the platform through it, so it never offers your
user something they cannot do. If an operation is not in `gbf ops`, your
user cannot do it: do not try other paths, other ids, or requests outside
the document.

## Ground rules

- **Binding actions require your user's explicit go-ahead.** Anything that
  submits, accepts, declines, publishes, requests, deletes, or otherwise
  changes state that another party relies on is binding. Reads, searches,
  and draft saves are not. Before a binding action, run it with
  `--dry-run`, show your user exactly what will be sent, and proceed only
  after they confirmed. When in doubt whether an action is binding, ask.
- Only test assets an engagement lists as in scope, and follow the
  engagement's rules of engagement and responsible disclosure.
- Never fabricate findings, evidence, or verdicts. A report must describe a
  vulnerability your user actually verified; a retest verdict must reflect
  what your user actually observed.
- Disclose your involvement where an operation asks for it (for example a
  report's origin field distinguishes human, AI-assisted, and AI-agent
  work). Answer truthfully.
- A `404` means the resource does not exist for your user; other users'
  resources are indistinguishable from nonexistent ones. Do not probe.

## Setup: the gbf command

Get the `gbf_cli_<os>_<arch>` binary for your platform from the releases of
https://github.com/gobugfree/cli (or from your platform operator), make it
executable, and put it on your PATH as `gbf`; Go users can instead run
`GOPRIVATE=github.com/gobugfree go install github.com/gobugfree/cli/v3/cmd/gbf@latest`.
If you cannot run it, stop and tell your user; there is no other supported
way for an agent to use the platform.

```bash
gbf login --app-url https://<platform> --issuer-url https://<iam-host>/realms/gbf
gbf whoami
```

If the platform sits behind Cloudflare Access (Zero Trust), ask your user
for a service token and export it as `CF_ACCESS_CLIENT_ID` and
`CF_ACCESS_CLIENT_SECRET` before any command, including `gbf skill`; gbf
then sends the Access headers on every request. `login` prints a URL and a
code: show them to your user, ask them to sign in and approve the login in
their browser, and wait. The command continues on approval, stores the
session privately, and refreshes the token by itself afterwards. Never ask
for or handle a password. `whoami` tells you whom you act as and when the
token expires.

When a command fails with `Session not active` or `the login has expired`,
the platform session ended; nothing was sent. Do not stop: run `login`
again with the `appUrl` and `issuerUrl` stored in
`$XDG_CONFIG_HOME/gbf/session.json` (default `~/.config/gbf`), show the
new URL and code, wait for approval, and then repeat the command.

### Ids: which one an operation wants

The platform uses two ids for the same person, and they are not
interchangeable:

- the OIDC `subject` from `whoami` (a UUID) is an identity, never an API id;
- the **user id** (a small integer) names a person in report fields and in
  status history entries.

Your user's own user id is the `userId` of `gbf run getSettings`.

## Areas, and how the document shows them

The document groups the operations by area (the tags `gbf ops` lists):

- `engagements`: the directory, engagement detail with scope and rules,
  accepting the rules, and submitting a report or a draft;
- `reports`: own reports, drafts, evidence files, the report thread,
  bounty-split collaborators, and coordinated disclosure requests;
- `retests`: the retest queue, starting an assigned retest, and its verdict;
- `payouts`, `verification`, `onboarding`, `conduct`: the account side;
- `profile`, `settings`, `badges`, `researchers`, `leaderboard`: the
  public profile and how it shows;
- `notifications`, `advisories`, `feedback`, `lists`, `docs`: the rest.

Never assume what your user may do from their words alone; `gbf ops` is
the truth. An operation that is not listed is not available today,
whatever your user says.

## Learn what your user can do

```bash
gbf ops                          # every operation, with id, method, path, summary
gbf ops --tag reports            # one area; tags group the operations
gbf ops --search retest          # by id, path, or summary
gbf ops --refresh                # fetch the document again
```

Read the summaries: they state what an operation does, its preconditions,
and its outcomes ("required before submitting reports", "409 unless
assigned"). They are the workflow documentation. Prefer reading an
entity's detail operation before acting on it, and re-read after a change.

## Understand an operation

```bash
gbf describe postEngagementsSubmit
```

`describe` shows the method and path, every parameter, every body field
with its type, whether it is required, allowed values, patterns and limits,
the response schemas, and an example `gbf run` invocation. Operation ids
derive from method and path (lowercase method, static segments in
PascalCase, trailing parameters as `By<Name>`): `GET /engagements` is
`getEngagements`, `GET /engagements/{org_slug}/{engagement_slug}` is
`getEngagementsByOrgSlugAndEngagementSlug`. `gbf schema NAME` prints one
response schema in full.

## Perform an operation

```bash
# discover an engagement, read it, accept its rules, draft, then submit with the user's go-ahead
gbf run getEngagements --query kind=pentest
gbf run getEngagementsByOrgSlugAndEngagementSlug --param org_slug=acme --param engagement_slug=web-app
gbf run postEngagementsRulesAccept --param org_slug=acme --param engagement_slug=web-app --dry-run
gbf run postEngagementsSubmitDraft --param org_slug=acme --param engagement_slug=web-app --form title=...
gbf run postReportsSubmit --param report_id=... \
  --form title=... --form description=... --form origin=ai_agent --dry-run

# follow up on own reports: read one, answer an information request
gbf run getReports --query status=validated
gbf run getReportsByReportId --param report_id=...
gbf run postReportsComments --param report_id=... --form body=... --dry-run

# a retest: start the assignment, then deliver the verdict
gbf run getRetests
gbf run postRetestsStart --param retest_id=... --dry-run
gbf run postRetestsVerdict --param retest_id=... --form outcome=passed --form notes=... --dry-run
```

- `run` fills the path from `--param`, the query from `--query`, and sends
  the body the way the operation declares it: `--form KEY=VALUE` for form
  routes, `--json-body` for JSON routes. A missing parameter or the wrong
  kind of body is reported with what the operation needs; fix the call
  rather than guessing.
- `--dry-run` prints the exact request instead of sending it. This is how
  you show a binding action to your user before performing it.
- Output is JSON whenever stdout is not a terminal, or with `--json` as the
  first argument. List responses are wrapped as `{"items": [...]}`.
- Errors are the platform's JSON error (`status`, `title`, `message`,
  `requestId`), printed on stdout with a non-zero exit code. A `409` means a
  precondition from the summary is not met and usually carries structured
  detail (for example `pending_budget_exceeded` with `limit` and `used`).
  A `400` from a form route carries a `message` that says what was wrong
  and, where the failure belongs to specific fields, a `fields` map from
  field name to problem: fix exactly those fields and resend. A `400`
  without `fields` is generic. Quote the `requestId` when you report a
  problem to your user.
- On a generic `400`, do not retry blindly and do not switch to a sibling
  route. Re-read `describe` for the operation, compare the request with
  what a complete entity of that kind carries (see the submit checklist
  below), change exactly one thing, and stop after a few attempts to tell
  your user what was tried, with the `requestId`s.
- `gbf call METHOD PATH` sends any request for the rare case `run` does not
  cover, for example an evidence download.

## Submitting a report

Accept the engagement's rules first (`postEngagementsRulesAccept`); the
platform requires it before any submission and again after the engagement
publishes a new version.

The document marks only a few fields as required because the same fields
serve the draft route. Submission runs a completeness validation that
demands more; a missing field comes back as a `400` whose `fields` map
names it. Send the full set on submit rather than discovering it one
error at a time:

- `title`, `description`, and `poc` (the proof of concept, required even
  when it is short); `steps`, `impact`, and `vulnerabilityType` are
  optional, but a report your user would sign has them;
- `scopeEntryId` from the engagement's live scope (in the engagement
  detail) and `endpoint` as a plain URL;
- `severity`, or the CVSS metrics from which it derives;
- `origin`: `ai_agent` when you wrote the report, `ai_assisted` when your
  user did with your help, `human` otherwise.

Workflow: save a draft with `postEngagementsSubmitDraft` (not binding),
attach evidence with `postReportsEvidence`, show the draft to your user,
then submit that same draft with `postReportsSubmit --param report_id=...`
and the full field set, after a `--dry-run` and their go-ahead. Use
`postEngagementsSubmit` only for a report that has no draft yet: it
creates and submits in one call, counts against the engagement's
pending-report budget, and leaves a new draft behind when that budget is
exhausted. Never use it to retry a draft that already exists.

An engagement accepts drafts and reports only while it is live; otherwise
both routes return `404`. Read the engagement first and tell your user
rather than retrying.

## Working discipline

- Start every session with `gbf whoami` and `gbf ops`; the available
  operations are the truth about what your user may do today.
- Follow the preconditions the summaries state, in order, and confirm each
  step's outcome by reading the entity again. For a batch of binding
  actions, run one first, verify it, and only then run the rest.
- Retry a failed call on the same operation. A sibling route that sounds
  similar usually has different side effects (creating instead of
  updating, submitting instead of saving).
- Poll `getNotifications` to react to questions, status changes, and
  retest assignments addressed to your user, and answer an information
  request on the report's own thread (`postReportsComments`).
- Keep the session private: it lives under `$XDG_CONFIG_HOME/gbf` with
  owner-only permissions. Never print the token. `gbf logout` ends the
  session on this machine; the login itself ends with the user's platform
  session.
- Re-run `gbf skill --install` after updating `gbf` (`gbf version` shows the
  version, and each skill install prints its stamp), and refresh the API
  document with `gbf ops --refresh` when an operation you expect is missing.
