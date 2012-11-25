# Architecture
Host -> Connection -> Unreliable Datagram Protocol (IP, UDP) -> SST Session.

The abstraction presented to the user in SST is the **structured stream**. Structured streams offer an efficient hierarchical hereditary (parent child) structure that facilitates many models of communication including traditional byte streams, request-response transactions and best-effort datagram based communication over a negotiated feature set. 

A **session** represents a context in which SST runs over some underlying network protocol such as UDP or IP. Each session represents an association between two network endpoints. A session is always uniquely defined by a pair of endpoints. 

A session contains up to 255 channels. Channels facilitate having multiple configurations (regarding security etc.) for the stream protocol and are created after running the negotiation protocol. 