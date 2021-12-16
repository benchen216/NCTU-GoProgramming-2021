======================
Homework assignment #0
======================

This is a demonstration of GoProgramming 21AU homework.

To turn in your homework, you need **a PR** and **the PR URL**.

Your PR needs to be created against the ``lab0`` branch.  (Not ``master``!)  The
name of your branch needs to start with ``<github-user-name>-lab#``, e.g.,
``username-lab2-attempt1``.  You need to create a sub-directory using exactly
your GitHub username as your working directory
(``NCTU-GoProgramming-2021/lab0/$github-user-name/``).  The hierarchy should be:
- NCTU-GoProgramming-2021(<- repository root)
    - lab0
        - username (<- your working directory)
            - Your files

In **every commit** in your PR, you can only change the files in your working
directory.  You may not touch anything else.  Failure to follow the rule can
cost you points.

Please make sure your PR passes the Github Action CI, and is compatible with
the latest Cloud9 on AWS (it uses Ubuntu 18.04 LTS) in ``us-east-1``.  You are
not required to use the AMI for doing the homework, but the grader is.  If your
code fails to build or run on it, **you can lose all points**.

Everyone should write his/her own code.  It is OK to discuss, but there should
not be duplicated code.  If duplication is found, **all points** for the
duplicated part of the latter submitter may be deducted.

Question
========

1. Run ``helloword.go``.
2. Write ``username.go`` program and print out github_username

Grading guideline
=================

This homework assignment has 2 points.  The grader will
run the following commands:

```#bash

cd NCTU-GoProgramming-2021/lab0/username
bash ../validate.py | grep "GET POINT"
```

