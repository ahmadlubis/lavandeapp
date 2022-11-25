class UserClient
  include HTTParty
  base_uri "http://localhost:10000/v1/user"
  format :json
  attr_reader :response

  # Login
  # POST /v1/user/login
  def login(input) 
    self.class.post(
      "/login",
      headers: {
        "Content-Type" => "application/json"
      },
      body: {
        "email": input['email'],
        "password": input['password']
      }.to_json
    )
  end

  # Get user info
  # GET /v1/user/index
  def get(token)
    self.class.get(
      "/me",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % token
      }
    )
  end

  # Register user
  # POST /v1/user/index
  def register(user_data)
    self.class.post(
      "/register",
      headers: {
        "Content-Type" => "application/json",
      },
      body: user_data.to_json
    )
  end

  # Update user data
  # PATCH /v1/user/
  def update(user_data, token)
    self.class.patch(
      "/me",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % token
      },
      body: user_data.to_json
    )
  end
end