Feature: All nodes must have a parent node/node tree
  In order to emulate an XML hierarchy
  As a developer
  I need to have everything be nested

  Scenario: fail to interact with parent-less node
    Given I create a node
    Then I should fail the following step
    And I label the node "node"

  Scenario: add an orphan node to node tree and succeed to interact with it
    Given I create a node
    And I create a node tree
    And I add the node to the tree
    And I label the node "node"
    And I give the node a value of "Hello, World!"
    Then data from node "node" should be "Hello, World!"
