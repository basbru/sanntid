// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i = -1;
pthread_mutex_t lock;

// Note the return type: void*
void* adder(){
        for(int x = 0; x < 1000000; x++){
		pthread_mutex_lock(&lock);
                i++;
		pthread_mutex_unlock(&lock);
    }
    return NULL;
}
void* subtractor(){
        for(int x = 0; x < 1000000; x++){
		pthread_mutex_lock(&lock);
                i--;
		pthread_mutex_unlock(&lock);
    }
    return NULL;
}


int main(){
	pthread_mutex_init(&lock,NULL);
    pthread_t adder_thr;
    pthread_create(&adder_thr, NULL, adder, NULL); 
    pthread_t subtractor_thr;
    pthread_create(&subtractor_thr, NULL, subtractor, NULL);
    pthread_join(subtractor_thr, NULL);
    pthread_join(adder_thr, NULL);
    printf("Done: %i\n", i);
	pthread_mutex_destroy(&lock);
    return 0;
    
}
