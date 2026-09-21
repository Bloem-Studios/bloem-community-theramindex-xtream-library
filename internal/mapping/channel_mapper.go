package mapping

import (
	"strconv"

	"github.com/Bloem-Studios/bloem-community-theramindex-xtream-library/internal/model"
	"github.com/Bloem-Studios/bloem-community-theramindex-xtream-library/internal/upstream/m3u"
	"github.com/Bloem-Studios/bloem-community-theramindex-xtream-library/internal/upstream/xtream"
)

func MapXtreamChannel(stream xtream.LiveStream) model.Channel {
	return model.Channel{
		ID: model.StableChannelID(model.SourceModeXtream, model.ChannelIdentity{
			UpstreamID: strconv.FormatInt(stream.StreamID, 10),
			GuideID:    stream.EPGChannelID,
			Name:       stream.Name,
			LogoURL:    stream.StreamIcon,
			StreamURL:  strconv.FormatInt(stream.StreamID, 10),
		}),
		SourceID:    model.LiveTVSourceID,
		Name:        stream.Name,
		Number:      strconv.Itoa(stream.Num),
		GuideID:     stream.EPGChannelID,
		LogoURL:     stream.StreamIcon,
		CategoryID:  stream.CategoryID,
		Catchup:     stream.CatchupAvailable(),
		CatchupMins: stream.TVArchiveDurationMinutes,
	}
}

func MapM3UChannel(entry m3u.Entry) model.Channel {
	return model.Channel{
		ID: model.StableChannelID(model.SourceModeM3UXMLTV, model.ChannelIdentity{
			GuideID:   entry.GuideID,
			Name:      entry.Name,
			LogoURL:   entry.LogoURL,
			StreamURL: entry.StreamURL,
		}),
		SourceID:  model.LiveTVSourceID,
		Name:      entry.Name,
		GuideID:   entry.GuideID,
		LogoURL:   entry.LogoURL,
		StreamURL: entry.StreamURL,
	}
}
