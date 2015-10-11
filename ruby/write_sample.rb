# $ bundle exec ruby ruby/write_sample.rb

require_relative './person.pb.rb'

person = Person.new
person.id = 456
person.name = 'NAME from Ruby'
person.email = 'ruby@example.com'

person.serialize_to_file('../sample_output_data.pb')

#person.parse_from_file('../sample_output_data.pb')
#p person
