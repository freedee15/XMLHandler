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
    Then the output should be:
    """
        <OtherNode></OtherNode>
    """

  Scenario: convert tree with 3 layers and check for data
    Given I create a node tree
    And I create a child node of node tree labelled "NewNode"
    And I create a child node of node "NewNode" labelled "OtherNode"
    And I select child node "NewNode" of node tree
    And I select child node "OtherNode" of selected node
    And I set the selected node's data to "foo"
    When I convert the selected node to XML
    Then the output should be:
    """
        <OtherNode> foo </OtherNode>
    """

  Scenario: convert node with 1 child
    Given I create a node tree
    And I create a child node of node tree labelled "NewNode"
    And I create a child node of node "NewNode" labelled "OtherNode"
    And I select child node "NewNode" of node tree
    And I select child node "OtherNode" of selected node
    And I set the selected node's data to "foo"
    And I select child node "NewNode" of node tree
    When I convert the selected node to XML
    Then the output should be:
      """
        <NewNode>
          <OtherNode> foo </OtherNode>
        </NewNode>
      """

  Scenario: convert node tree
    Given I create a node tree labelled "QuizXML"
    And I create a child node of node tree labelled "NewNode"
    And I create a child node of node "NewNode" labelled "OtherNode"
    When I convert the node tree to XML
    Then the output should be:
    """
    <QuizXML>
      <NewNode>
        <OtherNode></OtherNode>
      </NewNode>
    </QuizXML>
    """