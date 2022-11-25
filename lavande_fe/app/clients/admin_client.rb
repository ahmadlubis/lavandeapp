class AdminClient
    include HTTParty
    base_uri "http://localhost:10000/v1/admin"
    format :json
    attr_reader :response
  
    # Users list
    # POST /v1/admin/users
    def get_users(token, search, page) 
      self.class.get(
        "/users",
        headers: {
          "Authorization" => "Bearer %s" % token
        },
        query: {
          "limit" => 5,
          "offset" => 5 * page,
          "name" => search
        }
      )
      # if @response.success?
      #   p @response.parsed_response['data']
      #   @response.parsed_response['data']
      # else
      #   nil
      # end
    end

    # Change user status
    # PATCH /v1/admin/users
    def change_status(token, status_data) 
      self.class.patch(
        "/users",
        headers: {
          "Content-Type" => "application/json",
          "Authorization" => "Bearer %s" % token
        },
        body: status_data.to_json
      )
      # if @response.success?
      #   p @response.parsed_response
      #   @response.parsed_response
      # else
      #   nil
      # end
    end

    # Units list
    # POST /v1/admin/units
    def get_units(token, unit_data, page) 
      query = {
        "limit" => 5,
        "offset" => 5 * page
      }.merge(unit_data)
      self.class.get(
        "/units",
        headers: {
          "Authorization" => "Bearer %s" % token
        },
        query: query
      )
      # if @response.success?
      #   p @response.parsed_response['data']
      #   @response.parsed_response['data']
      # else
      #   nil
      # end
    end
end