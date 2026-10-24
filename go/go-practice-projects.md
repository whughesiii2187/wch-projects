# Small Go Projects: Building Retention Through Repetition

Six small, standalone projects. Each one touches something `budgetcli` never did. The goal is to fill gaps and get more reps on pointers, error handling and loops in unfamiliar contexts.

Plan: finish each project, archive it, redo it later from scratch (glancing at the old code only for "oh yeah, that's the framework"), then repeat.

This document explains **the ask**: what to build, in what order, and what to watch for. Finding the Go packages and functions that do each step is your job, and that is where most of the learning happens.

Every project has:

- **Goal and sample run**: what the finished program does, with exact input and output.
- **Plan in plain English**: the steps, with no Go in them.
- **Build stages**: the order to build it in. Run it after every stage.
- **Test data**: what to feed it so you can check your answer.
- **Watch out for**: traps in the ask itself. These say what can go wrong, not how to fix it.
- **Done when**: a finish line, so a project can't sprawl.

---

## How to break down any project

1. **Write the finished result as a sample.** Exact input, exact output. If you can't write the sample, the ask isn't clear yet. Sharpen it before touching code.
2. **Write the steps in plain English**, as if telling a coworker how to do it by hand with pen and paper. No Go words yet.
3. **Give each step a Go noun.** For every step, ask what the data is (string, number, list, lookup table, struct), what the action is (loop, if, function call), and what can fail (anything that returns an `err`).
4. **Search each noun separately.** Don't search "how to build a log counter". Search "go read file line by line", then "go split string into words". One small question gets one clear answer.
5. **Build one stage, run it, print what you have, then add the next.** Most bugs are visible the moment you print.

### Where to look

| You need | Go here |
|---|---|
| What a function takes and returns | pkg.go.dev, search the package, read the signature first, then the example |
| How a language feature works (maps, slices, loops, errors) | gobyexample.com, then go.dev/tour |
| What an error means | Paste the exact error text into a search, minus your own file names |
| "Is there something in the standard library for this?" | pkg.go.dev, search by plain words ("hash", "walk directory") |

---

## 1. Password generator (DONE)

**Goal:** generate a random password of a given length from the terminal.

Sample run:
```
$ ./passgen -len 20
k7Tq9mWxR2vBn4LpJs8D
$ ./passgen -len 8
length must be at least 16
```

**Plan in plain English:**
1. Read the length (and an optional "include symbols" switch) from the command line.
2. Reject lengths that are too short.
3. Decide which characters are allowed.
4. Once per position, pick one random allowed character.
5. The first and last positions never get a symbol.
6. Print the result.

**Build stages:**
1. Read the options from the command line and just print what you received.
2. Add the length check. Confirm 15 is rejected and 16 passes.
3. Build the set of allowed characters and print it.
4. Pick one random character and print it. Run it several times and confirm it changes.
5. Repeat the pick once per position and print the whole password.
6. Add the symbol switch.
7. Add the rule that the first and last positions never get a symbol.

**Test data:** length 15 (rejected), 16, 64, with and without the symbol switch.

**Watch out for:**
- Command-line options don't always hand you the plain value you typed. Read the function's signature to see what it returns.
- The "no symbol first or last" rule means those two positions draw from a smaller set of characters than the rest. Anything that depends on the size of the set has to use the right one.

**Revisit notes (your final version):** you handled the first and last position by drawing from a smaller pool. When you redo this from scratch, state the position rule in one sentence before writing any code. A stretch goal is to add a switch that excludes look-alike characters such as `0`, `O`, `l` and `1`.

**Done when:** it rejects short lengths, honors the symbol switch, and never puts a symbol first or last across about 100 runs.

---

## 2. File-backed todo list (DONE)

**Goal:** a terminal todo list that remembers its items between runs, stored in a JSON file instead of a database.

Sample run:
```
$ ./todo
New item: buy milk
1. buy milk

$ ./todo
New item: call dentist
1. buy milk
2. call dentist
```

**Plan in plain English:**
1. If the todo file doesn't exist yet, create it.
2. Load whatever items are already in it.
3. Ask for a new item and add it to the list.
4. Save the whole list back to the file.
5. Print the list.

**Build stages:**
1. Hardcode a list of three items in the program, convert it to JSON text and print it.
2. Write that text to a file. Open the file in an editor and confirm it looks right.
3. Read the file back and convert the text into a list in memory. Print the list.
4. Handle the file not existing yet.
5. Handle the file existing but being empty.
6. Take a new item from the user and add it to the list.
7. Save the updated list and print it.

**Test data:** a missing file, an empty file, a file with three items.

**Watch out for:**
- A missing file and an empty file are two different situations, and each can break the loading step in its own way.
- Converting JSON text into a list needs somewhere to put the result. Think about when that place has to exist.

**Revisit notes (your final version):** your version works. When you redo it, try storing each item as a small record (the text plus a done flag) instead of plain text. A stretch goal is a way to mark an item done, and a way to delete one.

**Done when:** items survive a restart, and a missing or empty file doesn't crash it.

---

## 3. "Is it up?" website checker (DONE)

**Goal:** check a list of sites and report each one's status code and response time.

Sample run:
```
google.com      200   143ms
microsoft.com   200    98ms
reddit.com      200   210ms
uscis.org       could not reach site: (reason)
```

**Plan in plain English:**
1. Have a list of site names.
2. For each one, note the time, request the page, and note how long it took.
3. If the request couldn't be made at all, say so and move on to the next site.
4. Otherwise, record the site, its status code and the time taken.
5. Print the results.

**Build stages:**
1. Request one hardcoded site and print its status code.
2. Add a time limit so a dead site can't hang the program.
3. Add timing and print how long the request took.
4. Loop over a list of sites and print each result.
5. Store each result as a record (site, code, time) and print them all at the end.
6. Handle a site that can't be reached, and confirm the others still run.

**Test data:** three real sites, one deliberately bad name, and (if you can find one) a site that returns a 404.

**Watch out for:**
- "The server answered with an error page" (like a 404) and "the request never succeeded" (like a bad name) are different situations. Decide how your program treats each one.
- Different network failures have different causes (a name that doesn't exist, a DNS server that misbehaves, a timeout). Read the error text.
- Something you open inside a loop and plan to close later may stay open longer than you expect. Check when "later" actually happens.

**Revisit notes (your final version):** your version keeps failed sites out of the results. When you redo it, decide on purpose whether a failed site should appear in the report as "down". A stretch goal is to take the site list from the command line instead of hardcoding it, and to print how many sites were up and how many were down.

**Done when:** a bad site doesn't stop the others, and every good site shows a code and a time.

---

## 4. Log file word/line counter (NOT STARTED)

**Goal:** read a text file, count how many times each word appears, and print the 10 most common.

**What "count" means here.** You are keeping a score sheet. Each word you meet gets a row, and the number beside it is how many times you've met that word so far. The number is never a line number and never a position in the file. It is just "times seen".

Worked example. The file contains:
```
error disk full
warning disk slow
error network down
```
Counting words, the finished score sheet is:
```
error    2
disk     2
full     1
warning  1
slow     1
network  1
down     1
```
Sample output (top 3):
```
error   2
disk    2
full    1
```

You can check this by hand, which is why it's the test data.

**Why words and not whole lines.** On a real log, every full line is different because of timestamps, so counting whole lines gives you 1 for nearly everything. Words like `error` and `timeout` repeat constantly, and that is where a useful answer is. The thing you count is the only difference between a word counter, a line counter and a pattern counter. Start with words. Once that works, you can try counting whole lines with the timestamp removed.

**Plan in plain English:**
1. Open the file.
2. Go through it one line at a time.
3. Break each line into words.
4. For each word, find its row on the score sheet and add one. If it has no row yet, it starts at one.
5. When the file is finished, the score sheet is complete, but it has no particular order.
6. Put the rows in order, highest count first.
7. Print the first 10.

**Build stages:**
1. Open the file and print every line. This proves the file reading works.
2. Print every word on its own line. This proves the splitting works.
3. Build the score sheet and print all of it. Compare it to the table above.
4. Put the rows into an ordered list and print it. It will still be in the wrong order.
5. Sort the list by count and print it.
6. Limit the output to the top 10.
7. Take the filename from the command line instead of hardcoding it.

**Test data:** the three-line example above. Then an empty file, a missing file, and a real log from your machine.

**Watch out for:**
- Punctuation stays attached to words unless you remove it, so `error:` and `error` count as different words. Decide whether to handle that now or in a later stage.
- `Error` and `error` are the same word to a human. Decide whether they should be the same to your program.
- When two words have the same count, their order is not guaranteed. Decide what the tie-break should be.
- A score sheet in Go has no built-in order. Look up what that means for sorting before you reach stage 5.
- Reading a file line by line has a limit on how long a line can be. Know what happens when a line is too long.

**Done when:** the three-line example matches the table above exactly, and a real log produces a believable top 10.

---

## 5. Duplicate file finder (NOT STARTED)

**Goal:** scan a folder and everything under it, and report files whose contents are identical.

Sample output:
```
Duplicate group (3 files):
  /photos/a/img1.jpg
  /photos/backup/img1.jpg
  /photos/old/copy of img1.jpg

Duplicate group (2 files):
  /docs/report.pdf
  /docs/archive/report-final.pdf
```

**Plan in plain English:**
1. Start at a folder you give it and visit every file under it, including in subfolders.
2. For each file, produce a short fingerprint of its contents. Two files with the same contents get the same fingerprint, and a different file almost certainly gets a different one.
3. Keep a list of which files share each fingerprint.
4. Any fingerprint shared by more than one file is a duplicate group.
5. Print the groups.
6. Never delete anything. This program only reports.

**Order of work:** get the whole thing working one file at a time first. Only after it works, make it handle several files at once. The speed-up is the point of the project, but it's much easier once you know the one-at-a-time version is right.

**Build stages (one at a time):**
1. Visit every path under a small test folder and print each one.
2. Print only files, not folders.
3. Produce a fingerprint for each file and print `path fingerprint`.
4. Group paths by fingerprint.
5. Print only the groups with more than one file.
6. Handle folders you can't read without the whole scan crashing.

**Warm-up before the several-at-once stage.** Don't change the real program yet. First write a tiny separate program that does something harmless several times at once: for example, five tasks that each add an entry to one shared list, and the program waits for all five to finish before printing. Run it with Go's race detector turned on and watch it complain. Then make the tasks take turns using the shared list, and confirm the complaint goes away. Only then apply the same idea to the real program.

**Build stages (several at once):**
7. Hash several files at the same time instead of one after another.
8. Make sure the program waits for all of them before printing the groups.
9. Make sure they don't corrupt the shared list of groups while writing to it.
10. Run it with the race detector on and confirm it reports nothing.

**Test data:** build a throwaway folder yourself: three different files, plus copies of two of them in subfolders, plus one empty file. You should get exactly one group for each file you copied.

**Watch out for:**
- Two empty files have identical contents. Decide whether empty files should count as duplicates.
- Same filename does not mean same contents, and different filenames can have identical contents. Only the contents decide.
- Some folders will refuse to be read. Decide what the program should do and say when that happens.
- On a folder with thousands of files, doing everything at once can run into limits on how many files can be open. Know that this can happen.
- Anything started inside a loop and sharing the loop's variable may not see the value you expect. Check what each task actually sees.

**Done when:** your throwaway folder produces exactly the groups you built into it, in both the one-at-a-time and the several-at-once versions, and the race detector is clean.

---

## 6. Write real tests for one of the above (NOT STARTED)

**Goal:** a test file that proves the core logic of an earlier project works, and runs with a single command.

**What this is.** A test is a small program that calls one of your functions with an input, checks the result against the answer you expected, and reports pass or fail. You write down a list of "given this, expect that" cases once, and from then on you never have to check by hand again.

**Which project to test.** The password generator or the log counter, since neither needs a network or a database. For the test to be possible, the logic must be a function that takes input and returns output, rather than being buried in `main`. Pulling it out into its own function is part of the project.

Sample cases (log counter):

| Input | Expected result |
|---|---|
| empty text | no words counted |
| `error` | error: 1 |
| `error error` | error: 2 |
| `Error error` | error: 2 (case doesn't matter) |
| `error   disk` | error: 1, disk: 1 (extra spaces don't matter) |

Sample cases (password generator):

| Input | Expected result |
|---|---|
| length 15 | rejected |
| length 16 | accepted, 16 characters long |
| any valid length, many runs | first and last characters are never symbols |

**Plan in plain English:**
1. Pick one function to test.
2. Write down the table of inputs and expected results.
3. Write the test so it goes through the table and checks each row.
4. Run it and see it pass.
5. Break your function on purpose and confirm the test fails. A test that has never failed has not proven anything.

**Build stages:**
1. Move the logic out of `main` into its own function, and confirm the program still works.
2. Write one test with one case and run it.
3. Add rows until the table has three or four cases.
4. Add the awkward cases: empty input, odd spacing, boundaries like length 15 and 16.
5. Break the function on purpose and confirm the test fails.
6. Fix it and confirm the test passes again.

**Watch out for:**
- Random output can never be compared to one fixed answer. Think about what you can check about it instead (its length, which characters it contains).
- A function that reads from the keyboard or prints directly to the screen is hard to test. That difficulty is a sign the logic is still mixed up with input and output.

**Revisit notes:** the "symbol as the last character" bug in the password generator is exactly what one row in this table would have caught immediately, without rerunning by hand.

**Done when:** the tests pass, a deliberate bug makes them fail, and adding a new case means adding one row.

---

## How to use this list

- The projects roughly build on each other, but go in whatever order keeps you moving.
- Each is small enough to **finish**. If one starts sprawling, cut scope.
- When one is done, archive it. When you redo it, start from an empty file and open the old code only to check the framework.
- No need to touch `budgetcli` until these feel done, or until you want to come back to it with fresh energy.
