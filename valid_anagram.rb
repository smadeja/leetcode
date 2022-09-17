# frozen_string_literal: true

def is_anagram(s, t)
  letter_count = Hash.new(0)

  s.each_codepoint do |letter|
    letter_count[letter] += 1
  end

  t.each_codepoint do |letter|
    new_letter_count = letter_count[letter] - 1

    return false if new_letter_count < 0

    letter_count[letter] = new_letter_count
  end

  letter_count.all? { |_key, value| value == 0 }
end
