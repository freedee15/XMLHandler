Feature: Access child node data easily
  In order to emulate an XML hierarchy
  As a developer
  I need to provide methods to interact with child nodes

  Scenario: access child node by it's label
    Given I create a node tree
    And I create a child node with the label "NewNode"
    And I give child node "NewNode" a value of "Hello, World!"
    When I retrieve data from child node "NewNode"
    Then the retrieved node data should be "Hello, World!"

  Scenario: fail to get non-existent child node
    Given I create a node tree
    And the next step should fail
    When I retrieve data from child node "NewNode"
    Then I should get the error 'no child with label "NewNode"'

  Scenario: Access all children
    Given I create a node tree
    And I create a child node with the label "NewNode"
    And I create a child node with the label "OtherNode"
    And I create a child node in "NewNode" with the label "NewNode1"
    And I create a child node in "NewNode" with the label "NewNode2"
    And I create a child node in "NewNode" with the label "NewNode3"
    Then the children array of the node tree should have 2 members
    And the children array of node "NewNode" should have 3 members