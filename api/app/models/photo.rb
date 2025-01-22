class Photo < ApplicationRecord
  has_and_belongs_to_many :albums
  has_and_belongs_to_many :categories

  validates :path, presence: true
  validates :filename, presence: true
end 