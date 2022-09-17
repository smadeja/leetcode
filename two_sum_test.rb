# frozen_string_literal: true

require "set"
require "minitest/autorun"
require_relative "two_sum"

class TestTwoSum < Minitest::Test
  TestCase = Struct.new(:nums, :target, :expected, keyword_init: true)

  test_cases = [
    TestCase.new(
      nums: [2, 7, 11, 15],
      target: 9,
      expected: [0, 1],
    ),
    TestCase.new(
      nums: [3, 2, 4],
      target: 6,
      expected: [1, 2],
    ),
    TestCase.new(
      nums: [3, 3],
      target: 6,
      expected: [0, 1],
    ),
  ]

  test_cases.each_with_index do |test_case, i|
    define_method "test_case_#{i+1}" do
      actual = two_sum(test_case.nums, test_case.target)
      assert_same_elements test_case.expected, actual
    end
  end

  def assert_same_elements(expected, actual)
    assert_equal frequency(expected), frequency(actual)
  end

  def frequency(array)
    result = Hash.new(0)
    array.each { |x| result[x] += 1 }

    result
  end
end
