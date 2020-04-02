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
    <?xml version="1.0" encoding="UTF-8" ?>
    <XMLTree>
      <NewNode>
        <OtherNode> foo </OtherNode>
      </NewNode>
    </XMLTree>
    """

  Scenario: import a simple xml tree from file
    Given I import the file "features/test.xml"
    When I select child node "NewNode" of node tree
    And I keep the selected node's data
    And I select child node "OtherNode" of node tree
    And I keep the selected node's data
    And I select child node "OtherNode1" of selected node
    And I keep the selected node's data
    Then the kept data should be "foobarbaz"


  Scenario Outline: test multiple invalid xml files to check for errors
    Given the next step should fail
    When I import the file "features/<filename>.xml"
    Then I should get the error <error>

    Examples:
    | filename         | error                      |
    | brokenprolog     | "invalid xml version"      |
    | dataintree       | "data in tree"             |
    | closewithoutopen | "node closed without open" |