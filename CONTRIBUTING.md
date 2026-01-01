# CONTRIBUTING V1

### Commit Messages Should follow the structure below:

```bash
<type>(<scope>): <short, imperative summary>

<body - optional>
```

---

### Use only these `type`:

- `feat – new feature`

- `fix – bug fix`

- `test – tests added/updated`

- `refactor – internal restructuring (no behavior change)`

- `docs – README, comments, etc`

- `build – CI, Docker, scripts`

That’s enough and Don’t invent new ones

---

Using `scope` is recommended

- `auth`
- `project`
- `resource`
- `http`
- `db`
- `middleware`
- `shutdown`
- `core`

---

### Examples (GOOD commits)

Auth Feature:

```bash
feat(auth): add password hashing and validation
```

Domain logic:

```bash
feat(project): enforce project ownership invariant
```

Transactions:

```bash
feat(resource): make resource creation transactional
```

Testing:

```bash
test(resource): add transaction rollback tests
```

Refactor:

```bash
refactor(http): extract error mapping into helper
```

Docs:

```bash
docs: add architecture overview to README
```
