# frozen_string_literal: true

require "set"
require "minitest/autorun"
require_relative "group_anagrams"

module ArrayRefinements
  refine Array do
    def deep_clone
      map(&:deep_clone)
    end
  end

  refine String do
    alias_method :deep_clone, :clone
  end
end

class TestGroupAnagrams < Minitest::Test
  TestCase = Struct.new(:strs, :expected, keyword_init: true)

  test_cases = [
    TestCase.new(
      strs: ["eat", "tea", "tan", "ate", "nat", "bat"],
      expected: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
    ),
    TestCase.new(
      strs: [""],
      expected: [[""]],
    ),
    TestCase.new(
      strs: ["a"],
      expected: [["a"]],
    ),
  ]

  test_cases.each_with_index do |test_case, i|
    define_method("test_case_#{i+1}") do
      actual = group_anagrams(test_case.strs)
      assert_deep_same_elements(test_case.expected, actual)
    end
  end

  using ArrayRefinements

  def assert_deep_same_elements(expected, actual)
    expected = expected.deep_clone
    expected.map!(&:sort!)
    expected = expected.to_set

    actual = actual.deep_clone
    actual.map!(&:sort!)
    actual = actual.to_set

    assert_equal expected, actual
  end
end
