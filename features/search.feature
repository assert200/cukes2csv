Feature: Search
  As a user
  I want search results
  So I can see the world

  Scenario: Search results
    * I search with:
      | search term | results 1 | results 2 | results 3 |
      | colours     | blue      | green     | red       |
      | music       | jazz      | pop       | rnb       |

  Scenario Outline: Search results count
    Given I on the search screen
    When I search for <search term> isn't searchable
    Then I get <count> search results
    Examples:
      | search term | count |
      | the         | 1000  |
      | big         | 20    |
      | test        | 3     |


