# frozen_string_literal: true

require "minitest/autorun"
require_relative "valid_anagram"

class TestIsAnagram < Minitest::Test
  TestCase = Struct.new(:s, :t, :expected, keyword_init: true)

  test_cases = [
    TestCase.new(
      s: "anagram",
      t: "nagaram",
      expected: true,
    ),
    TestCase.new(
      s: "rat",
      t: "car",
      expected: false,
    )
  ]

  test_cases.each_with_index do |test_case, i|
    define_method "test_case_#{i+1}" do
      actual = is_anagram(test_case.s, test_case.t)
      assert_equal test_case.expected, actual
    end
  end
end
