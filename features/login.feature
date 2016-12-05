Feature: Login
  As a user
  I want to login
  So I am authorised to use the application

  @wip
  Scenario: Login with correct credentials
    Given I on the login screen
    When I login with correct credentials
    Then I successfully login

  Scenario: Login with incorrect credentials
    Given I on the login screen
    When I login with incorrect credentials
    Then I am shown an invalid credentials error message
    And I do not login
    
