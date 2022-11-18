# app/models/product.rb
class User
    include ActiveModel::Model
    include ActiveModel::Attributes
    # attribute :id
    # attribute :email
    # attribute :password
    # attribute :full_name
    attr_accessor :id, :name, :nik, :email, :phone_no, :role, :status, :religion, :created_at, :updated_at
  
    def persisted?
      true
    end
  end