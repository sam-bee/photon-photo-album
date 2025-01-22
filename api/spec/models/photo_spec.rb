require 'rails_helper'

RSpec.describe Photo, type: :model do
  it { should have_and_belong_to_many(:albums) }
  it { should have_and_belong_to_many(:categories) }
  it { should validate_presence_of(:path) }
  it { should validate_presence_of(:filename) }
end 