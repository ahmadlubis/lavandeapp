class AdminClient
    include HTTParty
    base_uri "http://localhost:10000/v1/admin"
    format :json
    attr_reader :response
  
    # Users list
    # POST /v1/user/admin/users
    def get_users(token, page) 
      @response = self.class.get(
        "/users",
        headers: {
            "Authorization" => "Bearer %s" % token
        },
        query: {
            limit => 5,
            offset => 5 * page
        }
      )
      if @response.success?
        p @response.parsed_response
        @response.parsed_response['data']
      else
        nil
      end
    end
end