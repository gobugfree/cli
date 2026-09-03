---
name: gbf
description: Work the GObugfree bug bounty platform on behalf of your user through the gbf command. Use when asked to do anything on GObugfree (engagements, reports, triage, retests, payouts, organization or backoffice work) or to call its API. Covers the device-flow login, discovering the operations the user may call, and performing them safely.
---

# GObugfree Agent Skill

You are an AI agent acting on behalf of a human user of GObugfree, a bug
bounty platform. Your user may be a security researcher, a member of an
organization that runs engagements, or platform staff. You authenticate as
your user, you have exactly their permissions and nothing more, and every
action is recorded under their identity.

This file ships inside `gbf` and in its repository
(https://github.com/gobugfree/cli). Install it where your agent loads skills
from, for example `~/.claude/skills/gbf/SKILL.md` for Claude Code, and run
the same command again after updating `gbf`:

```bash
gbf skill --install
```

## The one rule that shapes everything: access is dynamic

The API is the same set of routes the web UI uses, and what your user may
call depends on their roles and memberships. The platform tells you exactly
that: its OpenAPI document is scoped to the logged-in user and lists only
the operations that user can perform, with parameters, fields, constraints,
and response shapes. There is no fixed list of routes to learn. You work
through the `gbf` command, which reads that document after login and drives
the platform through it, so it never offers your user something they cannot
do. If an operation is not in `gbf ops`, your user cannot do it: do not try
other paths, other ids, or requests outside the document.

## Ground rules

- **Binding actions require your user's explicit go-ahead.** Anything that
  submits, accepts, approves, rejects, publishes, pays, assigns, deletes, or
  otherwise changes state that another party relies on is binding. Reads,
  searches, and draft saves are not. Before a binding action, run it with
  `--dry-run`, show your user exactly what will be sent, and proceed only
  after they confirmed. When in doubt whether an action is binding, ask.
- Only test assets an engagement lists as in scope, and follow the
  engagement's rules of engagement and responsible disclosure.
- Never fabricate findings, evidence, or decisions. A report must describe a
  vulnerability your user actually verified; a triage decision must reflect
  what your user decided.
- Disclose your involvement where an operation asks for it (for example a
  report's origin field distinguishes human, AI-assisted, and AI-agent
  work). Answer truthfully.
- A `404` means the resource does not exist for your user; resources of
  other users or organizations are indistinguishable from nonexistent
  ones. Do not probe.

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
their browser, and wait. The command continues on
approval, stores the session privately, and refreshes the token by itself
afterwards. Never ask for or handle a password. `whoami` tells you whom you
act as and which groups (areas) they have, which is your first hint of the
role.

When a command fails with `Session not active` or `the login has expired`,
the platform session ended; nothing was sent. Do not stop: run `login`
again with the `appUrl` and `issuerUrl` stored in
`$XDG_CONFIG_HOME/gbf/session.json` (default `~/.config/gbf`), show the
new URL and code, wait for approval, and then repeat the command.

### Ids: which one an operation wants

The platform uses several ids for the same person, and they are not
interchangeable:

- the OIDC `subject` from `whoami` (a UUID) is an identity, never an API id;
- the **user id** (a small integer) names a person in report fields such as
  `assignee`, and in the audit log as `actorId`;
- the **member id** from an organization's members list names the
  membership, not the person, and is a different number.

When a field says "user id", pass the user id. Your user's own id is the
`userId` of `gbf run getSettings`. For assigning a report, read the
report's organization detail first: its `assignmentOptions` list every
active member with the id the assign operation takes. The members list
(`getOrganizationsMembers`, admins only) carries the same value as
`userId`; its `id` is the membership, never pass that as a user id.

## Roles, and how the document shows them

The same document serves three kinds of users; you recognize the role by
the areas (tags) `gbf ops` lists, and a user may combine them:

- **Researchers** work on the root routes: `engagements` (directory,
  detail, rules of engagement, submissions), `reports` (own reports,
  drafts, evidence, collaborators), `retests`, `payouts`, `notifications`,
  `profile`, `settings`, `badges`.
- **Organization members** see the `organizations` tag: their engagements
  and scope, the reports submitted to them (`getOrganizationsReports`,
  `getOrganizationsEngagementsReports`), and the decisions on them
  (accept, reject, request a retest, assign an internal owner).
- **Operations staff** see the `backoffice` tag: the triage queue
  (`postBackofficeTodosPickup`, `getBackofficeTodosByReportId`,
  `postBackofficeTodosValidate`, `postBackofficeTodosRequestInfo`, bounty
  proposals), engagement reviews, budgets, conduct cases, and badges.

Never assume a role from the user's words alone; `gbf ops` is the truth. A
user without the `organizations` tag is not an organization member today,
whatever they say, and the same goes for the backoffice.

## Learn what your user can do

```bash
gbf ops                          # every operation, with id, method, path, summary
gbf ops --tag organizations      # one area; tags group the operations
gbf ops --search report          # by id, path, or summary
gbf ops --refresh                # fetch the document again (after a role change)
```

Read the summaries: they state what an operation does, its preconditions,
and its outcomes ("required before submitting reports", "409 unless
active"). They are the workflow documentation. Prefer reading an entity's
detail operation before acting on it, and re-read after a change.

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
# a researcher: discover, read, draft, then submit with the user's go-ahead
gbf run getEngagements --query kind=pentest
gbf run getEngagementsByOrgSlugAndEngagementSlug --param org_slug=acme --param engagement_slug=web-app
gbf run postEngagementsSubmitDraft --param org_slug=acme --param engagement_slug=web-app --form title=...
gbf run postEngagementsSubmit --param org_slug=acme --param engagement_slug=web-app \
  --form title=... --form description=... --form origin=ai_agent --dry-run

# an organization member: review submitted reports, then decide
gbf run getOrganizationsReports --param org_slug=acme --query status=validated
gbf run getOrganizationsEngagementsReportsByReportId --param org_slug=acme \
  --param engagement_slug=web-app --param report_id=...
gbf run postOrganizationsEngagementsReportsAccept --param org_slug=acme \
  --param engagement_slug=web-app --param report_id=... --dry-run

# operations staff: pick up a report for triage and validate it
gbf run postBackofficeTodosPickup --param report_id=...
gbf run getBackofficeTodosByReportId --param report_id=...
gbf run postBackofficeTodosValidate --param report_id=... --form ... --dry-run
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
  cover, for example a download.

## Submitting a report

The document marks only a few fields as required because the same fields
serve the draft route. Submission runs a completeness validation that
demands more; a missing field comes back as a `400` whose `fields` map
names it. Send the full set on submit rather than discovering it one
error at a time:

- `title`, `description`, and `poc` (the proof of concept, required even
  when it is short); `steps`, `impact`, and `vulnerabilityType` are
  optional, but a report your user would sign has them;
- `scopeEntryId` from the engagement's live scope (`gbf run
  getOrganizationsScopesEngagementsByEngagementSlug` for members, the
  engagement detail for researchers) and `endpoint` as a plain URL;
- `severity`, or the CVSS metrics from which it derives;
- `origin`: `ai_agent` when you wrote the report, `ai_assisted` when your
  user did with your help, `human` otherwise.

Workflow: save a draft with `postEngagementsSubmitDraft` (not binding),
show the draft to your user, then submit that same draft with
`postReportsSubmit --param report_id=...` and the full field set, after a
`--dry-run` and their go-ahead. Use `postEngagementsSubmit` only for a
report that has no draft yet: it creates and submits in one call, counts
against the engagement's pending-report budget, and leaves a new draft
behind when that budget is exhausted. Never use it to retry a draft that
already exists.

An engagement accepts drafts and reports only once it is live: while its
only version is a draft, both routes return `404`. Read the engagement
first and tell your user rather than retrying. Bringing an engagement
live is the organization's job: add at least one in-scope entry to the
draft version, send it to review with
`postOrganizationsEngagementsSubmit` (a `400` says what the draft still
lacks), and wait for the backoffice to approve and publish it.

## Working discipline

- Start every session with `gbf whoami` and `gbf ops`; the available
  operations are the truth about your user's role today.
- Follow the preconditions the summaries state, in order, and confirm each
  step's outcome by reading the entity again. For a batch of binding
  actions, run one first, verify it, and only then run the rest.
- Retry a failed call on the same operation. A sibling route that sounds
  similar usually has different side effects (creating instead of
  updating, submitting instead of saving).
- Reports may carry a free-text `assetIdentifier` rather than a scope
  entry; when your user asks for per-asset analysis, check that field
  varies across reports before promising asset-level results, and say so
  when it does not.
- If your user has a notifications operation (`getNotifications`), poll it
  to react to questions, status changes, and assignments addressed to them.
- Keep the session private: it lives under `$XDG_CONFIG_HOME/gbf` with
  owner-only permissions. Never print the token. `gbf logout` ends the
  session on this machine; the login itself ends with the user's platform
  session.
- Re-run `gbf skill --install` after updating `gbf` (`gbf version` shows the
  version, and each skill install prints its stamp), and refresh the API
  document with `gbf ops --refresh` after a role or membership change.
