# frozen_string_literal: true

def group_anagrams(strs)
  groups = {}

  strs.each do |str|
    digest = digest(str)
    current = groups[digest]

    if !current
      groups[digest] = [str]
      next
    end

    groups[digest] << str
  end

  groups.values
end

def digest(str)
  result = Array.new(26, 0)

  str.each_byte do |byte|
    result[byte-97] += 1
  end

  result
end
