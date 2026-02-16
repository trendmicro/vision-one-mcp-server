## Title

[Feature/Fix/Refactor/Docs] Short summary of the change

---

## Description

Describe what this PR does and why it’s needed. Include any relevant context or issues.

> Example:
> This PR adds user authentication using JWT. It addresses part of the login flow for Issue #123.

---

## Changes

List the changes made in this PR:

- Added `/login` endpoint and JWT middleware
- Updated user model to store hashed passwords
- Created unit tests for the authentication flow

---

## How to Test

Steps to manually test this PR:

1. `git checkout <branch-name>`
2. `make mcpserver`
3. Move the `v1-mcp-server` binary onto `$PATH`
4. Configure the server with your AI tooling (differs depending on tool used)
5. _Describe specific test steps here..._

---

## Screenshots / Demos (if applicable)

_Add screenshots, UI recordings, or demos here._

---

## Related Issues

Closes #___
Related to #___

---

## Checklist

- [ ] I have tested my changes locally
- [ ] I have added or updated tests
- [ ] I have updated relevant documentation
- [ ] I have linked relevant issues
- [ ] I have assigned reviewers
