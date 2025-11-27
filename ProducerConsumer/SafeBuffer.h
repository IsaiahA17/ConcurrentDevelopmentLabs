#include <queue>
#include "Event.h"
#include "Semaphore.h"
/* SafeBuffer.h --- 
 * 
 * Filename: SafeBuffer.h
 * Description: 
 * Author: Joseph
 * Maintainer: 
 * Created: Tue Jan  8 12:30:23 2019 (+0000)
 * Version: 
 * Package-Requires: ()
 * Last-Updated: Tue Jan  8 12:30:25 2019 (+0000)
 *           By: Joseph
 *     Update #: 1
 * URL: 
 * Doc URL: 
 * Keywords: 
 * Compatibility: 
 * 
 */

/* Commentary: 
 * Must be thread safe using mutexes and semaphores
 * Must be of a queue structure
 * 
 */

/* Change Log:
 * 
 * 
 */

/* This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or (at
 * your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with GNU Emacs.  If not, see <http://www.gnu.org/licenses/>.
 */

/* Code: */

queue<Event> events;
std::mutex mutexLock;
Semaphore semaphore;

void put(Event event){
    //Need Semaphore in event that queue has a max capacity. i.e if statement that blocks threads with semaphore if full to prevent from placing in buffer 
    //and signalling if semaphore is up and and the queue size is less than the hypothetical size limit. Example code below for if statement
    //if(events.size() == maxSize) {
    //semaphore.wait()
    //}
    //else if(events.size() < maxSize && semaphore.wait() == true){
    //semaphore.signal()
    //}
    //Mutex lock to ensure only one thread at a time pushes an event, preventing multiple threads from pushing in potentially the same index in queue in event the push takes place at the same time
    mutexLock.lock()
    events.push(e);
    mutexLock.unlock()
}

void get(){
    //Mutex lock to ensure only one thread at a time pushes an event, preventing multiple threads from pushing in potentially the same index in queue in event the push takes place at the same time
    if(events.empty()) {
    semaphore.wait()
    }
    else if(!events.empty() && semaphore.wait() == true){
    semaphore.signal()
    }
    mutexLock.lock()
    events.pop();
    mutexLock.unlock()
}

/* SafeBuffer.h ends here */
