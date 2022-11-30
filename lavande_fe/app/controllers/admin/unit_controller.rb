class Admin::UnitController < ApplicationController

  def unit
    unit_input = params
    if !unit_input[:tower].nil? && !unit_input[:floor].nil? && !unit_input[:unit_no].nil?
      @units ||= session[:units]
      result = AdminClient.new.get_units(@token, unit_input.permit(:tower, :floor, :unit_no), 0)
      if result.success?
        @cur_unit = result.parsed_response['data'][0]
      else
        err_msg = result.parsed_response['error_message']
        redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when fetching units: %s" % err_msg
      end
    elsif !unit_input[:tower].nil? && !unit_input[:floor].nil?
      result = AdminClient.new.get_units(@token, unit_input.permit(:tower, :floor), 0)
      if result.success?
        @units ||= []
        for unit in result.parsed_response['data'] do
          @units << [unit['gov_id'], unit['unit_no'].to_i]
        end
        session[:units] = @units
      else
        err_msg = result.parsed_response['error_message']
        redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when fetching units: %s" % err_msg
      end
    end
  end

  def new
    unless session[:unit_data].nil?
      session[:unit_data].as_json().each do |name, value|
        params[name] = value
      end
      session.delete(:unit_data)
    end
  end

  def create
    unit_data = create_unit_params
    session[:unit_data] = unit_data
    result = AdminClient.new.create_unit(@token, unit_data)
    if result.success?
      redirect_to admin_unit_index_path, notice: "Successfully created unit %s" % unit_data[:gov_id]
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when creating unit: %s" % err_msg
    end
  end

  def edit
    result = AdminClient.new.edit(@token)
    if result.success?
      unit_data = result.parsed_response
      @unit = unit.new(unit_data)
    else
      redirect_to admin_unit_index_path, alert: "An error occurred when retrieving unit data"
    end
  end

  def update
    unit_data = unit_edit_params
    if !unit_data[:password].blank?
      # Check password
      if unit_data[:password] != unit_data[:password_confirmation]
        redirect_back fallback_location: unit_index_path, alert: "Passwords are not equal"
        return
      end
    else
      unit_data.delete(:password)
    end
    result = AdminClient.new.update(unit_data.except(:password_confirmation), @token)
    if result.success?
      redirect_to unit_index_path, notice: "Successfully updated unit data"
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when updating unit data: %s" % err_msg
    end
  end

  private

  # def unit_params
  #   params.require([:tower, :floor])
  #   params.permit(:tower, :floor)
  # end

  def create_unit_params
    params.require([:gov_id, :tower, :floor, :unit_no])
    params.permit(:gov_id, :tower, :floor, :unit_no)
  end
end
