// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

#include "Barrier.h"
#include "Event.h"
#include "SafeBuffer.h"
#include <iostream>
#include <thread>
#include <vector>


static const int num_threads = 100;
const int size=20;
SafeBuffer sBuffer

/*! \fn producer
    \brief Creates events and adds them to buffer
*/

void producer(std::shared_ptr<SafeBuffer<std::shared_ptr<Event>> theBuffer, int numLoops){

  for(int i=0;i<numLoops;++i){
    //Produce event and add to buffer
    std::shared_ptr<Event> e= std::make_shared<Event>(i); //Creates event
    theBuffer.put(e); //Multiple threads adding to the same buffer after creating an event may cause problems 
  }
  

}

/*! \fn consumer
    \brief Takes events from buffer and consumes them
*/

void consumer(std::shared_ptr<SafeBuffer<std::shared_ptr Event>> theBuffer, int numLoops){

  for(int i=0;i<numLoops;++i){
    //Produce event and add to buffer
    std::shared_ptr<Event> e= theBuffer->get();
    e->consume();  //Multiple threads consuming from the same buffer after creating an event may cause problems 
    // - consuming from an empty buffer causing need for mutex lock so one at a time and a semaphore that causes threads to wait if the buffer is empty and signals if the buffer is not empty
  }
  

}

int main(void){

  std::vector<std::thread> vt(num_threads);
  std::shared_ptr<SafeBuffer<std::shared_ptr<Event>> aBuffer( new Buffer<shared_ptr Event>(size));
  /**< Launch the threads  */
  int i=0;
  for(std::thread& t: vt){
    t=std::thread(updateTask,aBarrier,10);
  }
  /**< Join the threads with the main thread */
  for (auto& v :vt){
      v.join();
  }
  std::cout << sharedVariable << std::endl;
  return 0;
}
