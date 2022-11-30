class User
    include ActiveModel::Model
    include ActiveModel::Attributes

    attr_accessor :id, :name, :nik, :email, :phone_no, :role, :status, :religion, :created_at, :updated_at

    validates :name, presence: true
    validates :nik, presence: true
    validates :email, presence: true, length: {in:5..255}
    validates :phone_no, presence: true
    validates :religion, presence: true

    def initialize(attributes={})
      super
      @role ||= "resident"
      @status ||= "active"
    end

    def public_attributes
      {
        'name' => @name,
        'nik' => @nik,
        'email' => @email,
        'phone_no' => @phone_no,
        'religion' => @religion,
      }
    end
  
    def persisted?
      true
    end
  end