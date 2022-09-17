# frozen_string_literal: true

def two_sum(nums, target)
  num_indices = {}

  nums.each_with_index do |num, i|
    diff = target - num
    matching_index = num_indices[diff]

    return [i, matching_index] if matching_index

    num_indices[num] = i
  end

  []
end
