class UserClient
  include HTTParty
  base_uri "http://localhost:10000/v1/user"
  format :json
  attr_reader :response

  # Login
  # POST /v1/user/login
  def login(input) 
    @response = self.class.post(
      "/login",
      headers: {
        "Content-Type" => "application/json"
      },
      body: {
        "email": input['email'],
        "password": input['password']
      }.to_json
    )
    if @response.success?
      @response.parsed_response
    else
      nil
    end
  end

  # Get user info
  # GET /v1/user/index
  def get(token)
    @response = self.class.get(
      "/me",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % token
      }
    )
    # unless @response.parsed_response.key?("error_message")
    if @response.success?
      @response.parsed_response
    else
      nil
    end
  end

  # Register user
  # POST /v1/user/index
  def register(user_data)
    @response = self.class.post(
      "/register",
      headers: {
        "Content-Type" => "application/json",
      },
      body: user_data.to_json
    )
    # unless @response.parsed_response.key?("error_message")
    if @response.success?
      @response.parsed_response
    else
      nil
    end
  end

  # Update user data
  # PATCH /v1/user/
  def update(user_data, token)
    p user_data
    @response = self.class.patch(
      "/me",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % token
      },
      body: user_data.to_json
    )
    # unless @response.parsed_response.key?("error_message")
    if @response.success?
      @response.parsed_response
    else
      nil
    end
  end
end