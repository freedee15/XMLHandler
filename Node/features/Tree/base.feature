Feature: Create a base node that does not require a parent for other nodes to be housed in
  In order to require a parent for every node
  As a developer
  I must provide a base node

  Scenario: new node tree with node that identifies it as it's parent
    Given I create a node tree
    When I add a child node labelled "child" to the node tree
    Then node "child" should have the parent "root"