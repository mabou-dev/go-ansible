# go-ansible

## structure
```
cmd/
| ...
pkg/
| connection/
| | ...
| module/
| | ...
| playbook/
| | ...
```

instruction are written in a playbook
the library __pkg/playbook__ has the code refering to the playbooks.
Definition of the components, parsing, etc.

a connection is done on the client to execute task.
It can be local or though ssh.
the library __pkg/connection__ has the code fo the different communication methods.

a module is the tools used to execute the task
in the library __pkg/module__

## documentation

documentation is in folder `doc/`.
execute the command below to generate the UML diagrams with `plantUML`

```
plantuml -o doc/output doc/*
```

> __Note:__ `doc/output/*.png` files are set in gitignore to reduce size of git archive
