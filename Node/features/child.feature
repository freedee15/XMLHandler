Feature: Access child node data easily
  In order to emulate an XML hierarchy
  As a developer
  I need to provide methods to interact with child nodes

  Scenario: access child node by it's label
    Given I create a node
    And I create a child node with the label "NewNode"
    And I give child node "NewNode" a value of "Hello, World!"
    When I retrieve data from child node "NewNode"
    Then the retrieved node data should be "Hello, World!"

  Scenario: fail to get non-existent child node
    Given I create a node
    And I should fail the following step
    When I retrieve data from child node "NewNode"
    Then the retrieved node data should be ""