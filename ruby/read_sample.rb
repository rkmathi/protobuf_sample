# $ bundle exec ruby ruby/read_sample.rb

require_relative './person.pb.rb'

person = Person.new
person.parse_from_file('../sample_output_data.pb')
p person
