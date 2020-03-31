Feature: Provide a single struct-type to hold multiple kinds of data the same way
  In order to emulate an XML hierarchy
  As a developer
  I need to be able to interact with multiple types of data stored within the hierarchy

  Scenario: match node data types
    Given I create a node
    When I give the node a "int" to store
    Then I can retrieve the data from the node as a "int"