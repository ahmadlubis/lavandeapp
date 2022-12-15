class Admin::UnitController < ApplicationController
  require_relative '../../clients/admin/unit_client'
  PAGINATION_LIMIT = 10

  def index
    result = Admin::UnitClient.new(@token).index(unit_list_query)
      if result.success?
        @result = result.parsed_response
      else
        @err_msg = result.parsed_response['error_message']
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
    payload = params.permit(:gov_id, :tower, :floor, :unit_no)

    session[:unit_data] = payload
    result = Admin::UnitClient.new(@token).create(payload)
    if result.success?
      redirect_to admin_unit_index_path, notice: "Successfully created unit %s" % payload[:gov_id]
      session.delete(:unit_data)
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when creating unit: %s" % err_msg
    end
  end

  def edit
    result = Admin::UnitClient.new(@token).index(edit_unit_query)
    if result.success? && !result.parsed_response['data'].empty?
      unit_data = result.parsed_response['data'][0]
      @unit = Unit.new(unit_data)
    else
      redirect_to admin_unit_index_path, alert: "An error occurred when retrieving unit data"
    end
  end

  def update
    payload = params.require(:unit).permit(:id, :gov_id, :tower, :floor, :unit_no)
    payload[:id] = payload[:id].to_i

    result = Admin::UnitClient.new(@token).update(payload)
    if result.success?
      redirect_to admin_unit_index_path, notice: "Successfully updated unit %s data" % payload['gov_id']
    else
      err_msg = result.parsed_response['error_message']
      redirect_back fallback_location: admin_unit_index_path, alert: "An error occurred when updating unit data: %s" % err_msg
    end
  end

  private

  def unit_list_query
    params[:page] ||= 1
    params[:page] = params[:page].to_i
    query = params.permit(:page, :tower, :floor)
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = (query[:page] - 1) * PAGINATION_LIMIT
    query
  end

  def edit_unit_query
    params.require(:id)
    query = params.permit(:id)
    query[:page] = 1
    query[:limit] = PAGINATION_LIMIT
    query[:offset] = 0
    query
  end
end
