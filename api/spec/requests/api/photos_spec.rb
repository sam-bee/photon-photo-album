require 'rails_helper'

RSpec.describe "Api::Photos", type: :request do
  describe "GET /api/photos" do
    let!(:photo) { create(:photo) }
    let!(:category) { create(:category) }
    let!(:album) { create(:album) }

    before do
      photo.categories << category
      photo.albums << album
    end

    it "returns all photos" do
      get '/api/photos'
      expect(response).to have_http_status(:success)
      expect(json['data'].length).to eq(1)
    end

    it "filters by category" do
      get "/api/photos", params: { category: category.id }
      expect(json['data'].length).to eq(1)
    end

    it "filters by album" do
      get "/api/photos", params: { album: album.id }
      expect(json['data'].length).to eq(1)
    end
  end
end