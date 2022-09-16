# frozen_string_literal: true

require "minitest/autorun"
require_relative "contains_duplicate"

class TestContainsDuplicate < Minitest::Test
  TestCase = Struct.new(:nums, :expected, keyword_init: true)

  TEST_CASES = [
    TestCase.new(
      nums: [1, 2, 3, 1],
      expected: true
    ),
    TestCase.new(
      nums: [1, 2, 3, 4],
      expected: false
    ),
    TestCase.new(
      nums: [1, 1, 1, 3, 3, 4, 3, 2, 4, 2],
      expected: true
    )
  ]

  TEST_CASES.each_with_index do |test_case, i|
    define_method "test_case_#{i+1}" do
      actual = contains_duplicate(test_case.nums)
      assert_equal test_case.expected, actual
    end
  end
end
