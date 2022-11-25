class ApplicationController < ActionController::Base
    before_action :check_session

    private

    def check_session
        # Check for token in session and cookie
        # token = cookies[:access_token]
        token = session[:access_token]
        if token.nil?
            token = cookies[:access_token]
        end
        
        if token.nil?
            session[:is_logged_in] = false
            # flash[:alert] = "Please log in"
            redirect_to new_session_path
        else
            @token = token
            session[:is_logged_in] = true
        end
    end
end
