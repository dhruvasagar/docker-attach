# frozen_string_literal: true

require 'sinatra'

set :bind, '0.0.0.0'
set :port, 1234

get '/' do
  require 'pry'
  binding.pry
  'Dummy server, does nothing!'
end
