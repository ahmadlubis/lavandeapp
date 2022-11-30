class Unit
    include ActiveModel::Model
    include ActiveModel::Attributes

    attr_accessor :id, :gov_id, :tower, :floor, :unit_no, :ajb, :akte, :created_at, :updated_at

    validates :gov_id, presence: true
    validates :tower, presence: true
    validates :floor, presence: true
    validates :unit_no, presence: true

    def persisted?
      true
    end
  end