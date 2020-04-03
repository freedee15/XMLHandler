Feature: Convert individual node objects to XML output
  In order to output data to an XML file
  As a developer
  I need to convert each node to XML output

  Scenario: convert tree with 3 layers
    Given I create a node tree
    And I create a child node of node tree labelled "NewNode"
    And I create a child node of node "NewNode" labelled "OtherNode"
    And I select child node "NewNode" of node tree
    And I select child node "OtherNode" of selected node
    When I convert the selected node to XML
    Then The output should be "    <OtherNode></OtherNode>"

  Scenario: convert tree with 3 layers and check for data
    Given I create a node tree
    And I create a child node of node tree labelled "NewNode"
    And I create a child node of node "NewNode" labelled "OtherNode"
    And I select child node "NewNode" of node tree
    And I select child node "OtherNode" of selected node
    And I set the selected node's data to "foo"
    When I convert the selected node to XML
    Then The output should be "    <OtherNode>foo</OtherNode>"