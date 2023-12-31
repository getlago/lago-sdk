=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.45.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.0.1-SNAPSHOT

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for LagoAPI::ChargePropertiesGraduatedRangesInner
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe LagoAPI::ChargePropertiesGraduatedRangesInner do
  let(:instance) do
    LagoAPI::ChargePropertiesGraduatedRangesInner.new(
      from_value: 0,
      to_value: 10,
      flat_amount: '10',
      per_unit_amount: '0.5',
    )
  end

  describe 'test an instance of ChargePropertiesGraduatedRangesInner' do
    it 'should create an instance of ChargePropertiesGraduatedRangesInner' do
      expect(instance).to be_instance_of(LagoAPI::ChargePropertiesGraduatedRangesInner)
    end
  end
  describe 'test attribute "from_value"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  describe 'test attribute "to_value"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  describe 'test attribute "flat_amount"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  describe 'test attribute "per_unit_amount"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end
