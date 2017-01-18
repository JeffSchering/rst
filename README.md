# rst
Go script to adjust lengths of title underlines in reStructuredText files.

This script walks a directory tree, reads in each .rst file that it finds, 
and adjusts the title underline lengths if needed. It writes the changes back 
to the original files.

For example, this:
```
======
Administrator Tasks
======
```
becomes this:

```
===================
Administrator Tasks
===================
```
Similarly, this:
```
Administrator Tasks
--------
```
becomes this:
```
Administrator Tasks
-------------------
```

Lines that are longer than the title are not affected.

I did this on a Windows machine, so the line endings in the changed .rst files might be off.

## Usage

To **run the tests**, cd to your Go directory and type the following command:
```
go test github.com/JeffSchering/rst
```

To **install the script**, cd to your Go directory and type the following command:
```
go install github.com/JeffSchering/rst
```

To **run the script after installing**, go to your Go bin directory and type the following command:
```
rst.exe <filepath>
```
where \<filepath\> is the top level directory that contains your .rst files.
