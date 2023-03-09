import data.real.basic

section
variables a b c d : ℝ

#check (min_le_left a b : min a b ≤ a)
#check (min_le_right a b : min a b ≤ b)
#check (le_min : c ≤ a → c ≤ b → c ≤ min a b)

example : min a b = min b a :=
begin
  apply le_antisymm,
  { show min a b ≤ min b a,
    apply le_min,
    { apply min_le_right },
    apply min_le_left },
  { show min b a ≤ min a b,
    apply le_min,
    { apply min_le_right },
    apply min_le_left }
end

example : min a b = min b a :=
begin
  have h : ∀ x y, min x y ≤ min y x,
  { intros x y,
    apply le_min,
    apply min_le_right,
    apply min_le_left },
  apply le_antisymm, apply h, apply h
end

example : min a b = min b a :=
begin
  apply le_antisymm,
  repeat {
    apply le_min,
    apply min_le_right,
    apply min_le_left }
end

example : max a b = max b a :=
begin
  have h : ∀ x y, max x y ≤ max y x,
  {
    intros x y,
    apply max_le,
    apply le_max_right,
    apply le_max_left,
  },
  apply le_antisymm,
  apply h,
  apply h,
end

example : min (min a b) c = min a (min b c) := begin
  have x₀ : min (min a b) c ≤ a, sorry,
  have x₁ : min (min a b) c ≤ a, sorry,
  apply le_antisymm,
  { 
    show min (min a b) c ≤ min a (min b c),
    have h₀ : min (min a b) c ≤ min a b, by apply min_le_left,
    have h₁ : min a b ≤ a, by apply min_le_left,
    have h₂ : min (min a b) c ≤ a, begin
      apply le_trans h₀ h₁,
    end,

    have z₀ : min (min a b) c ≤ min a b, by apply min_le_left,
    have z₁ : min a b ≤ b, by apply min_le_right,
    have i₀ : min (min a b) c ≤ b, begin
      apply le_trans z₀ z₁,
    end,
    have i₁ : min (min a b) c ≤ c, by apply min_le_right,
    have i₂ : min (min a b) c ≤ min b c, begin
      apply le_min i₀ i₁,
    end,

    apply le_min h₂ i₂,
  },
sorry,
end

lemma aux : min a b + c ≤ min (a + c) (b + c) :=
sorry

example : min a b + c = min (a + c) (b + c) :=
sorry

#check (abs_add : ∀ a b : ℝ, abs (a + b) ≤ abs a + abs b)

example : abs a - abs b ≤ abs (a - b) :=
sorry

end

section
variables w x y z : ℕ

example (h₀ : x ∣ y) (h₁ : y ∣ z) : x ∣ z :=
dvd_trans h₀ h₁

example : x ∣ y * x * z :=
begin
  apply dvd_mul_of_dvd_left,
  apply dvd_mul_left
end

example : x ∣ x^2 :=
by apply dvd_mul_right

example (h : x ∣ w) : x ∣ y * (x * z) + x^2 + w^2 :=
sorry

end

section
variables m n : ℕ
open nat

#check (gcd_zero_right n : gcd n 0 = n)
#check (gcd_zero_left n  : gcd 0 n = n)
#check (lcm_zero_right n : lcm n 0 = 0)
#check (lcm_zero_left n  : lcm 0 n = 0)

example : gcd m n = gcd n m :=
sorry

end
