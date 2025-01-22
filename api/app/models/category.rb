class Category < ApplicationRecord
  has_and_belongs_to_many :photos

  validates :name, presence: true
end 