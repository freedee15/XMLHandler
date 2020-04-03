Feature: Export string data to an XML file
  In order to output data to an XML file
  As a developer
  I have to export it

  Scenario: Output a simple tree to a file
    Given I create a node tree labelled "XMLTree"
    And I create a child node of node tree labelled "NewNode"
    And I select child node "NewNode" of node tree
    And I create a child node of selected node labelled "OtherNode"
    And I select child node "OtherNode" of selected node
    And I set the selected node's data to "foo"
    When I export the node tree to "out.xml"
    Then the file "out.xml" should read:
    """
    <?xml version="1.0" encoding="UTF-8"?>
    <XMLTree>
      <NewNode>
        <OtherNode> foo </OtherNode>
      </NewNode>
    </XMLTree>
    """