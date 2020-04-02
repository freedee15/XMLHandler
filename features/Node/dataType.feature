Feature: Provide a single struct-type to hold multiple kinds of data the same way
  In order to emulate an XML hierarchy
  As a developer
  I need to be able to interact with multiple types of data stored within the hierarchy

  Scenario Outline: match node data types
    Given I create a node tree
    And I create a child node
    When I give the node a <type> to store
    And I retrieve the child node data as a <type>
    Then I should get no error

    Examples:
      | type     |
      | "bool"   |
      | "float"  |
      | "int"    |
      | "string" |

  Scenario Outline: fail to retrieve improper data type
    Given I create a node tree
    And I create a child node
    And I give the node a <type> to store
    And the next step should fail
    When I retrieve the child node data as a <otherType>
    Then I should get the error 'strconv.<error>: parsing "<parse>": invalid syntax'

    Examples:
      | type     | error     | parse | otherType |
      | "bool"   | Atoi      | true  | "int"     |
      | "float"  | Atoi      | 0.1   | "int"     |
      | "string" | ParseBool | s     | "bool"    |