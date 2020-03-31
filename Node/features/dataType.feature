Feature: Provide a single struct-type to hold multiple kinds of data the same way
  In order to emulate an XML hierarchy
  As a developer
  I need to be able to interact with multiple types of data stored within the hierarchy

  Scenario Outline: match node data types
    Given I create a node
    When I give the node a <type> to store
    Then I can retrieve the data from the node as a <type>

    Examples:
      | type     |
      | "bool"   |
      | "float"  |
      | "int"    |
      | "string" |

  Scenario Outline: fail to retrieve improper data type
    Given I create a node
    When I give the node a <type> to store
    Then I can't retrieve the data from the node as a <otherType>

    Examples:
      | type     | otherType |
      | "bool"   | "int"     |
      | "float"  | "int"     |
      | "string" | "bool"    |