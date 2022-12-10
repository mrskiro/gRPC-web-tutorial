import { PlainMessage } from "@bufbuild/protobuf"
import { useEffect, useState } from "react"
import { notificationClient } from "./connect"
import { Notification, NotificationType } from "./gen/proto/v1/notification_pb"

function App() {
  const [notifications, setNotifications] = useState<
    PlainMessage<Notification>[]
  >([])
  useEffect(() => {
    notificationClient.notificationList({ userId: "xxx" }).then((res) => {
      setNotifications(res.notifications)
    })
  }, [])
  return (
    <div>
      {notifications.map((v, i) => (
        <div key={i}>{JSON.stringify(v)}</div>
      ))}
      <button
        onClick={(e) => {
          e.preventDefault()
          notificationClient.sendNotification({
            userId: "xxx",
            type: NotificationType.EMAIL,
            body: "Emailです",
          })
        }}
      >
        send email!
      </button>
    </div>
  )
}

export default App
