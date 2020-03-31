# XMLHandler 

![Go](https://github.com/freedee15/XMLHandler/workflows/Go/badge.svg)

Version 0.1.0

```shell script
go get github.com/freedee15/XMLHandler@v0.1.0
```

---
_A Go package for interaction with XML files._

> package **_XMLHandler_**
>> package **_Node_**
>> - **Node**: A struct to hold data. Must have a parent to be usable.
>> - **Tree**: A struct to hold Nodes. Because it does not accept a parent it should be used as the base for a hierarchy.