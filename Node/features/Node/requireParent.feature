Feature: All nodes must have a parent node/node tree
  In order to emulate an XML hierarchy
  As a developer
  I need to have everything be nested

  Scenario: fail to interact with parent-less node
    Given I create a node
    And the next step should fail
    When I label the node "node"
    Then I should get the error "no parent"

  Scenario: add an orphan node to node tree and succeed to interact with it
    Given I create a node
    And I create a node tree
    And I add the node to the tree
    And I label the node "node"
    And I give the node a value of "Hello, World!"
    Then data from node "node" should be "Hello, World!"

  Scenario Outline: try to get data and make sure error is proper error
    Given I create a node
    And the next step should fail
    When I retrieve the node data as a <type>
    Then I should get the error "no parent"

    Examples:
      | type     |
      | "bool"   |
      | "float"  |
      | "int"    |
      | "string" |

  Scenario: try to get label to check error
    Given I create a node
    And the next step should fail
    When I get the node's label
    Then I should get the error "no parent"

  Scenario: try to get children array to check error
    Given I create a node
    And the next step should fail
    When I get the children array of the node
    Then I should get the error "no parent"

  Scenario: try to get parent to check error
    Given I create a node
    And the next step should fail
    When I get the parent of the node
    Then I should get the error "no parent"