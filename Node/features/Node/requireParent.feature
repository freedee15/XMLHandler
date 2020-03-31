Feature: All nodes must have a parent node/node tree
  In order to emulate an XML hierarchy
  As a developer
  I need to have everything be nested

  Scenario: fail to interact with parent-less node
    Given I create a node
    And I should fail the following step
    Then I label the node "node"