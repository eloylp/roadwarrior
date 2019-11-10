Feature: Central system
  In order to manage the program and orchestrate all the subsystems

  Background:
    Given that the central system is running

  Scenario: Central system lifecycle
    Given The central system is up
    And The central system has now 0 subsystems up
    When I start "echo" subsystem
    Then The central system has now 1 subsystems up
    When I stop "echo" subsystem
    Then The central system has now 0 subsystems up
    When I stop the central system
    Then The central system is down