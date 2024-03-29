module SessionsHelper
   # Returns true if the user is logged in, false otherwise.
   def logged_in
      session[:is_logged_in] == true
   end

   def get_user_id
      session[:user_id]
   end

   # Confirms a logged-in user.
   def logged_in_user
      unless logged_in
         redirect_to new_session_path
      end
   end

   def is_admin
      session[:role] == 'admin'
   end

   def is_owner
      session[:is_owner]
   end

end