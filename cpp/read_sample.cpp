// $ clang++ cpp/read_sample.cpp cpp/person.pb.cc -lprotobuf
// $ ./a.out

#include <iostream>
#include <fstream>
#include <string>

#include "person.pb.h"

int main(void) {
    GOOGLE_PROTOBUF_VERIFY_VERSION;

    Person person;
    std::ifstream fin("sample_output_data.pb", std::ios::binary);
    bool succeeded = person.ParseFromIstream(&fin);
    if ( !succeeded ) { std::cerr << "Failed to perse" << std::endl; }

    std::cout << "id:     " << person.id() << std::endl
              << "name:   " << person.name() << std::endl
              << "email:  " << person.email() << std::endl;

    return 0;
}
