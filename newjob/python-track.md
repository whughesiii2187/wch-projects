# Python Track: Relearning Python for Platform Automation

Seven lessons, P1 to P7, alternating weekly with the Terraform track. Planned at roughly 6 hours per week.

| Week | Lesson |
|---|---|
| 1 | P1: Core language refresh |
| 3 | P2: Files, errors, and generators |
| 5 | P3: Data modeling and type hints |
| 7 | P4: Project tooling and testing |
| 9 | P5: HTTP APIs |
| 10 | P6: Azure SDK |
| 11 | P7: Production-grade automation |

---

## How each lesson works

1. **Study.** Read the listed concepts in the official docs: the Python Tutorial and Library Reference at docs.python.org, and the Azure SDK for Python docs for P6.
2. **Explain it back.** Before writing any code, explain the "be able to explain" items to Claude in your own words. Expect follow-up questions.
3. **Build.** Attempt the exercise yourself.
4. **Review.** Bring your code and output to Claude for a PR-style review.
5. **Checkpoint.** Ask Claude to "quiz me on P#" before moving on.

**When you're stuck, hints escalate in this order:** the concept you're missing, then the relevant doc section, then the approach described in plain English. A small snippet comes last, and only for the specific piece you're stuck on. There are no finished solutions.

**Learning log:** keep one file per lesson recording what confused you, what fixed it, and anything that surprised you.

---

## P1: Core language refresh

**Be able to explain:**
- The differences between lists, tuples, dicts, and sets, and when to choose each.
- Comprehensions versus loops, and when a comprehension hurts readability.
- What `*args` and `**kwargs` do.
- Mutable versus immutable objects.
- Why a mutable default argument is a classic bug.

**Exercise:** use the provided `vm_inventory.csv`.
1. Find duplicate VMs.
2. Group VMs by resource group.
3. List VMs missing an `owner` tag.

**Rules:**
- Standard library only: the `csv` module, dicts, sets, and lists. No pandas.
- Before you write any code, write down your own definition of a "duplicate" and of a "missing owner."

**Done when:** your output matches what's actually planted in the file. Claude will tell you how many problems you caught and which kinds you missed, but not which rows.

**Log:** _

---

## P2: Files, errors, and generators

**Be able to explain:**
- What a context manager guarantees.
- The flow of `try` / `except` / `else` / `finally`.
- When to write a custom exception.
- How generators differ from lists in memory use.
- Why to prefer `pathlib` over `os.path`.

**Exercise:** a log analyzer that streams a large file, counts errors by category, and reports malformed lines without crashing.

**Done when:**
- It handles a file too large to load into memory.
- A missing file produces a clear message instead of a traceback.

**Log:** _

---

## P3: Data modeling and type hints

**Be able to explain:**
- When a `dataclass` beats a dict.
- When a plain function is better than a class.
- What type hints do and don't enforce at runtime.
- What the `ipaddress` module gives you.

**Exercise:** model a network definition (a VNet with subnets) loaded from YAML. Validate that every subnet falls inside the VNet range and that no subnets overlap.

**Done when:** each invalid case produces a specific error naming the offending subnet.

**Log:** _

---

## P4: Project tooling and testing

**Be able to explain:**
- What a virtual environment isolates.
- What `pyproject.toml` replaces.
- The difference between a linter, a formatter, and a type checker.
- What `pytest` fixtures and parametrized tests are for.

**Exercise:** package the P3 tool as an installable CLI, with `ruff`, `mypy`, and a test suite.

**Done when:**
- `pip install .` in a fresh virtual environment gives you a working command.
- All checks pass.
- Every validation error path has a test.

**Log:** _

---

## P5: HTTP APIs

**Be able to explain:**
- The HTTP status code classes.
- Which errors are worth retrying and which aren't.
- Exponential backoff.
- The common pagination patterns: link headers, cursors, and page numbers.
- Why you should always set timeouts.

**Exercise:** pull all repositories for a GitHub organization through the paginated API, handle rate limits, and export the results to CSV.

**Done when:** a simulated failure partway through doesn't lose completed work.

**Log:** _

---

## P6: Azure SDK

**Be able to explain:**
- The order in which `DefaultAzureCredential` tries credential sources.
- How the SDK handles paging.
- What a long-running operation (the `begin_` methods) is, and why it exists.
- The difference between management-plane and data-plane clients.

**Exercise:** a subscription hygiene report listing untagged resources, unattached managed disks, and unassociated public IPs.

**Done when:**
- It runs against your subscription with no stored secrets.
- The output is something you'd hand to a manager.

**Log:** _

---

## P7: Production-grade automation

**Be able to explain:**
- `logging` versus `print`, and how log levels work.
- Exit codes, and why CI depends on them.
- Configuration from environment variables versus command-line flags.
- CLI subcommands.

**Exercise:** harden the P6 tool so it can run unattended in a pipeline, failing with a non-zero exit code when it finds violations above a configurable threshold.

**Done when:** a teammate could use it correctly with nothing but `--help`.

**Log:** _

---

## Optional capstone (only if the team uses Python with Terraform)

Build a plan gate that parses `terraform show -json` output and blocks risky changes inside the T7 pipeline. It combines P7 and T7.
