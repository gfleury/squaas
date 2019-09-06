[![Build Status](https://travis-ci.org/gfleury/squaas.svg?branch=master)](https://travis-ci.org/gfleury/squaas) [![codecov](https://codecov.io/gh/gfleury/squaas/branch/master/graph/badge.svg)](https://codecov.io/gh/gfleury/squaas)

# SQUAAS

SQUAAS (SQL query as a service). Despite the hipster name, this is all a mistake.
An automated solution to a bad habit. However sometimes we need to look for a proper way of doing wrong things.

## Query Flow

```text
[On Hold] (Query can still be edited)
    |
[Ready] - After enough approvals -> [Approved] - Query is added to the running queue -> [Done/Failed]
```

## Goals

- Remove developers direct access to the databases (Only PostgreSQL support for now)
- Enforce some behaviors on the queries (must have transactions, no delete/update without where, hehe)
- Enforce queries to be run only with tickets (only Jira support for now)
- Enforce a flow where the queries have to be 'reviewed/approved' by other developers
- Prevent multiple 'user queries' from running in parallel in the database
- Simple 'built in' SQL parser (users can't add broken queries, uses https://github.com/xwb1989/sqlparser)
- Last but not least, learn some react
